<script lang="ts">
import Icon from '@iconify/svelte';
import { onMount } from 'svelte';

let darkMode = $state(false);
const toggleTheme = () => {
	const html = document.documentElement;
	if (darkMode) {
		html.style.colorScheme = 'light';
		document.cookie = 'theme=light;path=/;max-age=31536000';
	} else {
		html.style.colorScheme = 'dark';
		document.cookie = 'theme=dark;path=/;max-age=31536000';
	}
	darkMode = !darkMode;
};

onMount(() => {
	const getCookie = (name: string) => {
		const value = `; ${document.cookie}`;
		const parts = value.split(`; ${name}=`);
		if (parts.length === 2) return parts.pop()?.split(';').shift();
	};

	const savedTheme = getCookie('theme');
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
	transition-property: left;
	transition-duration: var(--transition-duration);
	&.checked {
		left: calc(var(--toggle-width) - var(--toggle-height) + (var(--indicator-padding) / 2));
	}
	& :global(svg) {
		padding: 0.1em;
	}
}

input[type='checkbox']:is(.toggle) {
	--toggle-height: 1.75em;
	--toggle-width: 3em;
	--indicator-padding: 0.5em;

	transition-property: all;
	position: relative;

	&::before {
		background-color: var(--background-neutral-alpha);
	}
	&::after {
		content: '';
		transition-property: left;
		transition-duration: var(--transition-duration);
		background-color: var(--inverse-text-color);
	}
	&:checked {
		&::after {
			content: '';
			left: calc(var(--toggle-width) - var(--toggle-height) + (var(--indicator-padding) / 2));
		}
	}
}
</style>
