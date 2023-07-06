export function getThemePreference() {
    const themeStorage = localStorage.getItem("color-theme");
    const isThemeStorage = "color-theme" in localStorage;
    const classList = document.documentElement.classList;
    const themePreference = window.matchMedia(
        "(prefers-color-scheme: dark)"
    ).matches;

    if (
        (isThemeStorage && themeStorage === "dark") ||
        (!isThemeStorage && themePreference)
    ) {
        classList.add("dark");
        localStorage.setItem("color-theme", "dark");
    } else {
        classList.remove("dark");
        localStorage.setItem("color-theme", "light");
    }
}