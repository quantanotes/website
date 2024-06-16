if (typeof window !== 'undefined') {
    var storedDarkMode = JSON.parse(localStorage.getItem('darkMode')) || false
    var darkMode = $state(storedDarkMode)

    if (storedDarkMode) {
        document.documentElement.classList.add('dark')
    }
}

export default {
    get darkMode() {
        return darkMode
    },

    toggle() {
        if (typeof window === 'undefined') return
        darkMode = !darkMode
        document.documentElement.classList.toggle('dark')
        localStorage.setItem('darkMode', JSON.stringify(darkMode))
    },
}
