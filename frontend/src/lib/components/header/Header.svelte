<script lang="ts">
import { resolve } from '$app/paths';
import Themetoggle from '$lib/components/theme/toggle.svelte';
import { setStylePropertyCallback } from '$lib/effects/setStylePropertyCallback.svelte';
import Icon from '@iconify/svelte';
import { swipeable } from '@svelte-put/swipeable';
import { slide } from 'svelte/transition';
import Modal from '../Modal.svelte';
import Search from './Search.svelte';
import User from './User.svelte';

let modal = $state<Modal>()!;
</script>

<header>
	<div>
		<a class="neutral" href={resolve('/')}><span>SteamInputDB.com</span></a>
	</div>
	<button
		onclick={() => {
			modal.setScrimOpacityMulti(1);
			modal.toggle();
		}}
		aria-label="Menu">
		<Icon icon="mdi:menu" width="100%" height="100%" />
	</button>
	<Search />
	<div>
		<User />
		<Themetoggle />
	</div>
</header>
<Modal bind:this={modal} --backdrop-filter="blur(2px)">
	<aside
		transition:slide={{ duration: 196, axis: 'x' }}
		use:swipeable={{
			direction: 'left'
		}}
		use:setStylePropertyCallback
		onsetstyleproperty={(e) => {
			const { name, value } = e.detail;
			if (name === '--swipe-distance-x') {
				let swipeDist = Number(value.replace('px', '').trim());
				swipeDist = Math.min(Math.max(-window.innerWidth / 0.75, swipeDist), 0);
				modal.setScrimOpacityMulti(1 - Math.abs(swipeDist) / (window.innerWidth / 0.75));
			}
		}}
		onswipeend={modal.swipeend}>
		<div>
			<div>
				<User />
			</div>
			<Themetoggle />
		</div>
	</aside>
</Modal>

<style lang="postcss">
header {
	padding: 1em;

	background: var(--card-background-noise);

	display: grid;
	grid-template-columns: minmax(1ch, auto) 1fr min-content;
	align-items: center;
	box-shadow: 0 0px 4px var(--shadow-color);
	gap: 1em;
	transition-property: all;
	& > :first-child {
		display: grid;
		grid-template-columns: minmax(1ch, min-content);
		& > * {
			display: unset;
			overflow: hidden;
			width: auto;
			text-overflow: ellipsis;
		}
	}

	& > :nth-child(2) {
		display: none;
	}

	& > :last-child {
		display: grid;
		grid-auto-flow: column;
		gap: 1em;
		margin-left: auto;
	}
}

@media (orientation: portrait) {
	header {
		& > :first-child {
			display: none;
		}
		& > :nth-child(2) {
			display: block;
		}
		& > :last-child {
			display: none;
		}
	}
}

.neutral {
	font-size: 1.4em;
	text-decoration: none;
	color: inherit;
	position: relative;
	width: min-content;
}

a {
	font-size: 1.4em;
	white-space: nowrap;
	display: grid;
	place-items: center;
	grid-auto-flow: column;
	gap: 0.25em;
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
	width: 2em;
	height: 2em;
}

aside {
	display: grid;
	min-width: 75%;
	width: fit-content;
	transform: translateX(clamp(-100%, var(--swipe-distance-x), 0px));
	& > div {
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
	background: var(--card-background-noise);
	backdrop-filter: blur(1px);
	height: 100%;
}
</style>
