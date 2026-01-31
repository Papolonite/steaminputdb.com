import { describe, expect, it, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import { page } from 'vitest/browser';
import '../../../css/main.pcss';
import { applyUserThemePreferences } from './applytheme.svelte';
import ThemeToggle from './toggle.svelte';

describe('ThemeToggle', () => {
    interface TestCase {
        name: string;
        existingCookie?: string;
        prefersColorScheme: 'dark' | 'light';
        action?: () => Promise<void>;
        expected: {
            darkMode: boolean;
            colorScheme: '' | 'light' | 'dark';
            cookieContains: string;
        };
    }

    const testCase: TestCase[] = [
        {
            name: 'SUCCESS_UseSavedDarkTheme_OverridesBrowserPreference',
            existingCookie: 'theme=dark',
            prefersColorScheme: 'light',
            expected: {
                darkMode: true,
                colorScheme: '',
                cookieContains: 'theme=dark'
            }
        },
        {
            name: 'SUCCESS_UseSavedLightTheme_OverridesBrowserPreference',
            existingCookie: 'theme=light',
            prefersColorScheme: 'dark',
            expected: {
                darkMode: false,
                colorScheme: '',
                cookieContains: 'theme=light'
            }
        },
        {
            name: 'SUCCESS_UseBrowserDarkPreference_WhenNoCookie',
            prefersColorScheme: 'dark',
            expected: {
                darkMode: true,
                colorScheme: '',
                cookieContains: 'theme=dark'
            }
        },
        {
            name: 'SUCCESS_UseBrowserLightPreference_WhenNoCookie',
            prefersColorScheme: 'light',
            expected: {
                darkMode: false,
                colorScheme: '',
                cookieContains: 'theme=light'
            }
        },
        {
            name: 'SUCCESS_ToggleDarkToLight',
            existingCookie: 'theme=dark',
            prefersColorScheme: 'dark',
            action: async () => {
                const checkboxEl = page.getByRole('checkbox', { name: 'Theme Toggle' });
                await checkboxEl.click();
            },
            expected: {
                darkMode: false,
                colorScheme: 'light',
                cookieContains: 'theme=light'
            }
        },
        {
            name: 'SUCCESS_ToggleLightToDark',
            existingCookie: 'theme=light',
            prefersColorScheme: 'light',
            action: async () => {
                const checkboxEl = page.getByRole('checkbox', { name: 'Theme Toggle' });
                await checkboxEl.click();
            },
            expected: {
                darkMode: true,
                colorScheme: 'dark',
                cookieContains: 'theme=dark'
            }
        },
        {
            name: 'SUCCESS_DefaultToLightTheme_WhenNoBrowserPreference',
            prefersColorScheme: 'light',
            expected: {
                darkMode: false,
                colorScheme: '',
                cookieContains: 'theme=light'
            }
        }
    ];

    testCase.forEach((tc) => {
        it(tc.name, async () => {
            document.cookie.split(';').forEach((c) => {
                document.cookie = c
                    .replace(/^ +/, '')
                    .replace(/=.*/, `=;expires=${new Date(0).toUTCString()};path=/`);
            });
            document.documentElement.style.colorScheme = '';
            if (tc.existingCookie) {
                document.cookie = `${tc.existingCookie};path=/`;
            }
            const matchMediaMock = vi.fn((query: string) => ({
                matches: query.includes('dark') ? tc.prefersColorScheme === 'dark' : false,
                media: query,
                onchange: null,
                addListener: vi.fn(),
                removeListener: vi.fn(),
                addEventListener: vi.fn(),
                removeEventListener: vi.fn(),
                dispatchEvent: vi.fn()
            }));
            vi.stubGlobal('matchMedia', matchMediaMock);

            render(ThemeToggle);

            await tc.action?.();
            const toggleInput = page.getByRole('checkbox', { name: 'Theme Toggle' });
            const checkbox = toggleInput.element() as HTMLInputElement;
            expect(checkbox.checked).toBe(tc.expected.darkMode);

            const colorScheme = document.documentElement.style.colorScheme;
            expect(colorScheme).toBe(tc.expected.colorScheme);

            expect(document.cookie).toContain(tc.expected.cookieContains);
            vi.unstubAllGlobals();
        });
    });
});

describe('applyUserThemePreferences', () => {
    interface TestCase {
        name: string;
        existingCookie?: string;
        prefersColorScheme: 'dark' | 'light';
        expected: {
            colorScheme: 'dark' | 'light';
            transitionDuration?: string;
        };
    }

    const testCases: TestCase[] = [
        {
            name: 'SUCCESS_SetsDark_WhenCookieDark',
            existingCookie: 'theme=dark',
            prefersColorScheme: 'light',
            expected: {
                colorScheme: 'dark'
            }
        },
        {
            name: 'SUCCESS_SetsLight_WhenCookieLight',
            existingCookie: 'theme=light',
            prefersColorScheme: 'dark',
            expected: {
                colorScheme: 'light'
            }
        },
        {
            name: 'SUCCESS_SetsDark_WhenPrefersDark_NoCookie',
            prefersColorScheme: 'dark',
            expected: {
                colorScheme: 'dark'
            }
        },
        {
            name: 'SUCCESS_SetsLight_WhenPrefersLight_NoCookie',
            prefersColorScheme: 'light',
            expected: {
                colorScheme: 'light'
            }
        },
        {
            name: 'SUCCESS_ApplyTransitionDuration_Callback',
            existingCookie: 'theme=light',
            prefersColorScheme: 'dark',
            expected: {
                colorScheme: 'light',
                transitionDuration: 'var(--transitionDuration)'
            }
        }
    ];

    testCases.forEach((tc) => {
        it(tc.name, async () => {
            document.cookie.split(';').forEach((c) => {
                document.cookie = c
                    .replace(/^ +/, '')
                    .replace(/=.*/, `=;expires=${new Date(0).toUTCString()};path=/`);
            });
            document.documentElement.style.colorScheme = '';
            document.documentElement.style.transitionDuration = '';
            document.body.style.transitionDuration = '';
            if (tc.existingCookie) {
                document.cookie = `${tc.existingCookie};path=/`;
            }
            const matchMediaMock = vi.fn((query: string) => ({
                matches: query.includes('dark') ? tc.prefersColorScheme === 'dark' : false,
                media: query,
                onchange: null,
                addListener: vi.fn(),
                removeListener: vi.fn(),
                addEventListener: vi.fn(),
                removeEventListener: vi.fn(),
                dispatchEvent: vi.fn()
            }));
            vi.stubGlobal('matchMedia', matchMediaMock);

            const applyTransition = applyUserThemePreferences();

            expect(document.documentElement.style.colorScheme).toBe(tc.expected.colorScheme);

            if (tc.expected.transitionDuration) {
                vi.useFakeTimers();
                applyTransition?.();
                vi.runAllTimers();
                expect(document.documentElement.style.transitionDuration).toBe(tc.expected.transitionDuration);
                expect(document.body.style.transitionDuration).toBe(tc.expected.transitionDuration);
                vi.useRealTimers();
            }

            vi.unstubAllGlobals();
        });
    });
});
