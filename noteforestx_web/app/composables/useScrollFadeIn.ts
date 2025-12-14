import { onMounted, onBeforeUnmount, nextTick } from 'vue'
import gsap from 'gsap'
import ScrollTrigger from 'gsap/ScrollTrigger'

let registered = false

export interface ScrollFadeInOptions {
    selector?: string
    y?: number
    x?: number
    duration?: number
    ease?: string
    stagger?: number
    start?: string
    toggleActions?: string
    direction?: 'up' | 'down' | 'left' | 'right'
}

export function useScrollFadeIn(options: ScrollFadeInOptions = {}) {
    const {
        selector = '.animate-card',
        y = 50,
        x = 50,
        duration = 1,
        ease = 'power2.out',
        stagger = 0.1,
        start = 'top 85%',
        toggleActions = 'play none none reverse',
        direction = 'up'
    } = options

    // ✅ 只记录当前 composable 创建的 triggers
    const triggers: ScrollTrigger[] = []

    onMounted(async () => {
        await nextTick()

        if (!registered) {
            gsap.registerPlugin(ScrollTrigger)
            registered = true
        }

        const elements = document.querySelectorAll<HTMLElement>(selector)

        elements.forEach((el, i) => {
            const fromVars: gsap.TweenVars = { autoAlpha: 0 }
            const toVars: gsap.TweenVars = {
                autoAlpha: 1,
                duration,
                ease,
                delay: i * stagger,
                scrollTrigger: {
                    trigger: el,
                    start,
                    toggleActions
                }
            }

            switch (direction) {
                case 'up':
                    fromVars.y = y
                    toVars.y = 0
                    break
                case 'down':
                    fromVars.y = -y
                    toVars.y = 0
                    break
                case 'left':
                    fromVars.x = x
                    toVars.x = 0
                    break
                case 'right':
                    fromVars.x = -x
                    toVars.x = 0
                    break
            }

            const tween = gsap.fromTo(el, fromVars, toVars)

            // ✅ 拿到这个 tween 对应的 ScrollTrigger
            if (tween.scrollTrigger) {
                triggers.push(tween.scrollTrigger)
            }
        })
    })

    onBeforeUnmount(() => {
        // ✅ 只 kill 自己的
        triggers.forEach(t => t.kill())
        triggers.length = 0
    })
}