import { onMounted, nextTick, onBeforeUnmount } from 'vue'
import gsap from 'gsap'
import ScrollTrigger from 'gsap/ScrollTrigger'

let registered = false

export interface ScrollFadeInOptions {
    selector?: string
    y?: number
    duration?: number
    ease?: string
    stagger?: number
    start?: string
    toggleActions?: string
}

/**
 * 通用滚动进入动画
 */
export function useScrollFadeIn(options: ScrollFadeInOptions = {}) {
    const {
        selector = '.animate-card',
        y = 50,
        duration = 1,
        ease = 'power2.out',
        stagger = 0.1,
        start = 'top 85%',
        toggleActions = 'play none none reverse'
    } = options

    onMounted(async () => {
        await nextTick()

        // ✅ 只注册一次
        if (!registered) {
            gsap.registerPlugin(ScrollTrigger)
            registered = true
        }

        const elements = document.querySelectorAll<HTMLElement>(selector)

        elements.forEach((el, i) => {
            gsap.fromTo(
                el,
                { autoAlpha: 0, y },
                {
                    autoAlpha: 1,
                    y: 0,
                    duration,
                    ease,
                    delay: i * stagger,
                    scrollTrigger: {
                        trigger: el,
                        start,
                        toggleActions
                    }
                }
            )
        })
    })

    // 可选：组件卸载时清理
    onBeforeUnmount(() => {
        ScrollTrigger.getAll().forEach(t => t.kill())
    })
}