import { browser } from '$app/environment';

const getCookie = (name: string) => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts?.length === 2) {
        return parts.pop()?.split(';').shift();
    }
};

export type ApplyTransitionDuration = () => void;

export const applyUserThemePreferences = (): ApplyTransitionDuration|undefined => {
    if (browser === false) {
        return;
    }
    const savedTheme = getCookie('theme');
    const prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    const isDark = savedTheme === 'dark' || (!savedTheme && prefersDark);
    const html = document.documentElement;
    html.style.colorScheme = isDark ? 'dark' : 'light';
    // prevent flashbang
    return () => {
        setTimeout(() => {
            [document.documentElement, document.body].forEach((e) => {
                e.style.transitionDuration = 'var(--transitionDuration)';
            });
        });
    };
};
