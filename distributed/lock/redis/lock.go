package redis

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-redis/redis/v8"
	"time"
)

type Locker struct {
	client          *redis.Client // Redis客户端
	script          *redis.Script // 解锁脚本
	ttl             time.Duration // 过期时间
	tryLockInterval time.Duration // 重新获取锁间隔
}

func NewLocker(client *redis.Client, ttl, tryLockInterval time.Duration) *Locker {
	return &Locker{
		client:          client,
		script:          redis.NewScript(unlockScript),
		ttl:             ttl,
		tryLockInterval: tryLockInterval,
	}
}

func (l *Locker) GetLock(resource string) *Lock {
	return &Lock{
		client:          l.client,
		script:          l.script,
		resource:        resource,
		randomValue:     gofakeit.UUID(),
		ttl:             l.ttl,
		tryLockInterval: l.tryLockInterval,
		watchDog:        make(chan struct{}),
	}
}

type Lock struct {
	client          *redis.Client // Redis客户端
	script          *redis.Script // 解锁脚本
	resource        string        // 锁定的资源
	randomValue     string        // 随机值
	ttl             time.Duration // 过期时间
	tryLockInterval time.Duration // 重新获取锁间隔
	watchDog        chan struct{} // 看门狗
}

func (l *Lock) Lock(ctx context.Context) error {
	err := l.TryLock(ctx)
	if err == nil {
		return nil
	}

	if !errors.Is(err, ErrLockFailed) {
		return err
	}

	ticker := time.NewTicker(l.tryLockInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ErrTimeout
		case <-ticker.C:
			err := l.TryLock(ctx)
			if err == nil {
				return nil
			}

			if !errors.Is(err, ErrLockFailed) {
				return err
			}
		}
	}
}

func (l *Lock) TryLock(ctx context.Context) error {
	success, err := l.client.SetNX(ctx, l.resource, l.randomValue, l.ttl).Result()
	if err != nil {
		return err
	}

	if !success {
		return ErrLockFailed
	}

	go l.startWatchDog()

	return nil
}

func (l *Lock) Unlock(ctx context.Context) error {
	err := l.script.Run(ctx, l.client, []string{l.resource}, l.randomValue).Err()
	close(l.watchDog)
	return err
}

func (l *Lock) startWatchDog() {
	ticker := time.NewTicker(l.ttl / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), l.ttl/3*2)
			ok, err := l.client.Expire(ctx, l.resource, l.ttl).Result()
			cancel()

			if err != nil || !ok {
				return
			}
		case <-l.watchDog:
			return
		}
	}
}
