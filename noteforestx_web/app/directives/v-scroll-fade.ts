import type { DirectiveBinding } from 'vue'
import gsap from 'gsap'
import ScrollTrigger from 'gsap/ScrollTrigger'

gsap.registerPlugin(ScrollTrigger)

interface ScrollFadeOptions {
    direction?: 'up' | 'down' | 'left' | 'right'
    x?: number
    y?: number
    duration?: number
    ease?: string
    delay?: number
}

export default {
    mounted(el: HTMLElement, binding: DirectiveBinding<ScrollFadeOptions>) {
        const opts = binding.value || {}
        const fromVars: gsap.TweenVars = { autoAlpha: 0 }
        const toVars: gsap.TweenVars = {
            autoAlpha: 1,
            duration: opts.duration ?? 1,
            ease: opts.ease ?? 'power2.out',
            delay: opts.delay ?? 0,
            scrollTrigger: {
                trigger: el,
                start: 'top 85%',
                toggleActions: 'play none none reverse'
            }
        }

        switch (opts.direction) {
            case 'up':
                fromVars.y = opts.y ?? 50
                toVars.y = 0
                break
            case 'down':
                fromVars.y = -(opts.y ?? 50)
                toVars.y = 0
                break
            case 'left':
                fromVars.x = opts.x ?? 50
                toVars.x = 0
                break
            case 'right':
                fromVars.x = -(opts.x ?? 50)
                toVars.x = 0
                break
        }

        const tween = gsap.fromTo(el, fromVars, toVars)
        ;(el as any)._scrollFadeTween = tween
    },
    unmounted(el: HTMLElement) {
        const tween = (el as any)._scrollFadeTween
        if (tween) tween.kill()
    }
}