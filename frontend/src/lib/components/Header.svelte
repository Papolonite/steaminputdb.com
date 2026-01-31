<script lang="ts">
import { resolve } from '$app/paths';
import Themetoggle from '$lib/components/theme/toggle.svelte';
</script>

<header>
	<a href={resolve('/')}><span>SteamInputDB.com</span></a>
	<Themetoggle />
</header>

<style lang="postcss">
header {
	padding: 1em;
	background: var(--card-color);
	display: grid;
	grid-template-columns: auto min-content;
	align-items: center;
	box-shadow: 0 0px 4px var(--shadow-color);
}

span {
	font-weight: bold;
	font-size: 1.4em;
}

a {
	text-decoration: none;
	color: inherit;
	position: relative;
	width: min-content;

	&::after {
		content: '';
		height: 3px;
		width: 100%;
		bottom: 0em;
		transform: translateY(0.25em);
		left: 0;
		position: absolute;
		background: var(--text-color);
		scale: 0 1;
		transition:
			scale var(--transition-duration) var(--scale-delay, 0ms) ease-in-out,
			translate calc(var(--transition-duration) * 2) var(--translate-delay, 0ms) ease-in-out;
	}
	&:hover,
	&:focus-within {
		&::after {
			scale: 1 1;
		}
	}
	&:hover {
		& + li::after {
			translate: -100%;
			--scale-delay: var(--transition-duration);
			--translate-delay: var(--transition-duration);
		}
	}
	&:has(+ :hover) {
		&::after {
			translate: 100%;
			--scale-delay: var(--transition-duration);
			--translate-delay: var(--transition-duration);
		}
	}
}
</style>
