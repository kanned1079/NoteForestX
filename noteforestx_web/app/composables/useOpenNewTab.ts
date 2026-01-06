export const openInNewTab = (url: string) => {
    if (import.meta.server) return

    const newWindow = window.open(url, '_blank')
    if (newWindow) {
        newWindow.focus()
    } else {
        console.warn('this operation not allowed.')
    }
}