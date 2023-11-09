package main

import "time"

// TODO: Rate limit by user

type CommandRateLimiter struct {
    LastUsed time.Time
    Cooldown time.Duration
}

func NewCommandRateLimiter(cooldown time.Duration) *CommandRateLimiter {
    return &CommandRateLimiter{
        Cooldown: cooldown,
    }
}

func (c *CommandRateLimiter) CanExecute() bool {
    return time.Since(c.LastUsed) >= c.Cooldown
}

func (c *CommandRateLimiter) Execute() {
    c.LastUsed = time.Now()
}

