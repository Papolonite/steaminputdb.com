<script lang="ts">
import { page } from '$app/state';
import Icon from '@iconify/svelte';
import { onMount } from 'svelte';

let darkMode = $state(page.data.theme == 'dark');
const toggleTheme = () => {
	const html = document.documentElement;
	const wrap = () => {
		// no-op, just for the transition
		if (darkMode) {
			html.style.colorScheme = 'light';
			document.documentElement.setAttribute('data-theme', 'light');
			document.cookie = 'theme=light;path=/;max-age=31536000';
		} else {
			html.style.colorScheme = 'dark';
			document.documentElement.setAttribute('data-theme', 'dark');
			document.cookie = 'theme=dark;path=/;max-age=31536000';
		}
		darkMode = !darkMode;
	};

	if (!document.startViewTransition) {
		wrap();
		return;
	}
	try {
		// firefox does not support view-transition OPTIONS yet...
		// it doesnt show the same transition artifacts as chrome for some reason
		// so no view-transition is actually fine!
		document.startViewTransition({
			types: ['theme'],
			update: wrap
		});
	} catch {
		wrap();
	}
};

onMount(() => {
	const savedTheme = page.data.theme;
	const prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
	const isDark = savedTheme === 'dark' || (!savedTheme && prefersDark);

	darkMode = isDark;

	if (isDark && !savedTheme) {
		document.cookie = 'theme=dark;path=/;max-age=31536000';
	} else if (!isDark && !savedTheme) {
		document.cookie = 'theme=light;path=/;max-age=31536000';
	}
});
</script>

<div class="wrapper">
	<input
		name="Theme toggle"
		type="checkbox"
		class="toggle"
		checked={darkMode}
		onchange={() => toggleTheme()}
		aria-label="Theme Toggle" />
	<div class="icon {darkMode ? 'checked' : ''}">
		{#if darkMode}
			<Icon icon="ph:moon" width="1.2em" height="1.2em" />
		{:else}
			<Icon icon="ph:sun" width="1.2em" height="1.2em" />
		{/if}
	</div>
</div>

<style lang="postcss">
.wrapper {
	position: relative;
	&::view-transition-old(theme, theme-icon) {
		animation-name: none !important;
	}

	&::view-transition-new(theme, theme-icon) {
		animation-name: none !important;
	}
}

.icon {
	--toggle-height: 1.75em;
	--toggle-width: 3em;
	--indicator-padding: 0.5em;
	width: auto;
	height: var(--toggle-height);
	position: absolute;
	top: 0;
	bottom: 0;
	display: grid;
	align-items: center;
	pointer-events: none;
	left: calc(var(--indicator-padding) / 2);
	&.checked {
		left: calc(var(--toggle-width) - var(--toggle-height) + (var(--indicator-padding) / 2));
	}
	& :global(svg) {
		padding: 0.1em;
	}
	transition: left var(--transition-duration) var(--default-ease);
	view-transition-name: theme-icon;
}

input[type='checkbox']:is(.toggle) {
	--toggle-height: 1.75em;
	--toggle-width: 3em;
	--indicator-padding: 0.5em;

	position: relative;

	&::before {
		background-color: var(--background-neutral-alpha);
	}
	&::after {
		content: '';
		background-color: var(--inverse-text-color);
		view-transition-name: theme;
	}
	&:checked {
		&::after {
			content: '';
			left: calc(var(--toggle-width) - var(--toggle-height) + (var(--indicator-padding) / 2));
		}
	}
}
</style>
