<script lang="ts">
import { resolve } from '$app/paths';
import { page } from '$app/state';
import SC2 from '$lib/assets/SC2_Alt.svg?component';
import { setStylePropertyCallback } from '$lib/attachments/setStylePropertyCallback.svelte';
import Themetoggle from '$lib/components/theme/toggle.svelte';
import Icon from '@iconify/svelte';
import { swipeable } from '@svelte-put/swipeable';
import { cubicIn, cubicOut } from 'svelte/easing';
import { fly } from 'svelte/transition';
import Modal from '../Modal.svelte';
import Search from './Search.svelte';
import User from './User.svelte';

let modal = $state<Modal>()!;
</script>

<header>
	<a class="home" href={resolve('/')}>
		<SC2 height="1.6em" />
		<span>SteamInputDB</span>
	</a>
	<button
		class="plain"
		onclick={() => {
			modal.setScrimOpacityMulti(1);
			modal.toggle();
		}}
		aria-label="Menu">
		<Icon icon="mdi:menu" width="100%" height="100%" />
	</button>
	{#if !page.url.pathname.endsWith('/search')}
		<Search />
	{:else}
		<div></div>
	{/if}
	<div>
		<User />
		<div>
			<Themetoggle />
		</div>
	</div>
</header>
<Modal bind:this={modal} --backdrop-filter="blur(2px)">
	<aside
		in:fly={{ duration: 196, x: '-100%', y: 0, easing: cubicOut, opacity: 1 }}
		out:fly={{ duration: 196, x: '-100%', y: 0, easing: cubicIn, opacity: 1 }}
		use:swipeable={{
			direction: 'left'
		}}
		use:setStylePropertyCallback
		onsetstyleproperty={(e) => {
			const { name, value } = e.detail;
			if (name === '--swipe-distance-x') {
				let swipeDist = Math.min(Number(value.replace('px', '').trim()), 0);
				modal.setScrimOpacityMulti(1 - Math.abs(swipeDist) / (window.innerWidth / 0.75));
			}
		}}
		onswipeend={modal.swipeend}>
		<div>
			<a class="home" href={resolve('/')}>
				<SC2 height="1.6em" />
				<span>SteamInputDB</span>
			</a>
			<Themetoggle />
		</div>
	</aside>
</Modal>

<style lang="postcss">
header {
	padding: 1em;
	position: relative;
	isolation: isolate;

	min-height: 6em;
	z-index: 2;

	&::after {
		content: '';
		position: absolute;
		inset: 0;
		z-index: -1;
		background: var(--card-background-noise);
		background-color: var(--card-color);
	}

	width: 100%;
	overflow-x: clip;

	display: grid;
	grid-template-columns: minmax(3.2em, 1fr) auto minmax(2em, 1fr);
	@media (orientation: portrait) {
		grid-template-columns: minmax(2em, 1fr) auto minmax(2em, 1fr);
	}
	align-items: center;
	box-shadow: 0 0px 4px var(--shadow-color);
	gap: 1em;
	transition-property: all;

	& > :first-child {
		margin-right: auto;
		padding: 0.25em;
	}

	& > :nth-child(2) {
		display: none;
	}

	& > :last-child {
		display: grid;
		grid-auto-flow: column;
		gap: 1em;
		margin-left: auto;
		place-items: center;
		overflow: clip;
		max-width: 100%;
		padding: 0.25em;
		overflow-clip-margin: 2em;
	}
}
.home {
	grid-template-columns: min-content auto;
	font-weight: bold;
	height: fit-content;
	& span {
		overflow: hidden;
		text-overflow: ellipsis;
		width: 100%;
	}
}

@media (orientation: portrait) {
	header {
		& > :first-child {
			display: none;
		}
		& > :nth-child(2) {
			display: grid;
		}
		& > :last-child {
			& > :last-child {
				display: none;
			}
		}
	}
}

a {
	font-size: 1.4em;
	white-space: nowrap;
	display: grid;
	place-items: center;
	grid-auto-flow: column;
	gap: 0.5em;
	color: var(--text-color);
	&:hover,
	&:focus,
	&:focus-within {
		color: var(--color-primary);
	}
}

button {
	display: grid;
	place-items: center;
	padding: 0;
	border: none;
	box-shadow: none;
	aspect-ratio: 1 / 1;
	max-width: 2em;
}

aside {
	display: grid;
	grid-template-rows: min-content;

	min-width: 75%;
	width: fit-content;
	height: 100%;

	background: var(--card-background-noise);
	backdrop-filter: blur(1px);
	background-color: light-dark(
		color-mix(in srgb, var(--card-color), transparent 80%),
		color-mix(in srgb, var(--card-color), transparent 20%)
	);

	/* &::before {
		content: '';
		position: absolute;
		inset: 0;
		background: var(--card-background-noise);
		z-index: -1;
	} */

	transition: background var(--transition-duration) var(--default-ease);
	transform: translateX(clamp(-100%, var(--swipe-distance-x), 0px));
	& > :first-child {
		display: grid;
		grid-template-columns: auto min-content;
		min-width: 20ch;
		padding: 1em;
		gap: 1em;
		width: 100%;
		& > :first-child {
			margin-right: auto;
		}
	}
}
</style>
