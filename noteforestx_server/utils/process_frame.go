package utils

import (
	"fmt"
	"time"
)

//func (this *Logger) DisplayLoadingAnime(message string) {
//	frames := []rune{'-', '\\', '|', '/'}
//	for {
//		for _, r := range frames {
//			fmt.Printf("\r%c %s", r, message)
//			time.Sleep(200 * time.Millisecond)
//		}
//	}
//}

func (l *Logger) DisplayLoadingAnime(message string) {
	if l.running {
		return // 避免重复启动
	}
	l.stopChan = make(chan struct{})
	l.running = true

	go func() {
		frames := []rune{'-', '\\', '|', '/'}
		start := time.Now()
		for {
			for _, r := range frames {
				select {
				case <-l.stopChan:
					fmt.Printf("\r✓ %s (Time: %s)\n", message, time.Since(start).Truncate(time.Millisecond))
					l.running = false
					return
				default:
					// 输出动画 + 已经过的时间
					elapsed := time.Since(start).Truncate(time.Second)
					fmt.Printf("\r%c [%s] %s", r, elapsed, message)
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()
}

func (l *Logger) StopLoadingAnime() {
	if l.running {
		close(l.stopChan)
	}
}
