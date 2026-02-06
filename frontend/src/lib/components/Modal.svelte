<script lang="ts">
import { beforeNavigate } from '$app/navigation';
import type { SwipeEndEventDetail } from '@svelte-put/swipeable';
import type { Snippet } from 'svelte';
import { fade } from 'svelte/transition';
let {
	opened = $bindable(false),
	slide_touch = $bindable(true),
	children
}: {
	opened?: boolean;
	slide_touch?: boolean;
	children?: Snippet;
	'--background'?: string;
	'--backdrop-filter'?: string;
	'--background-opacity'?: `${number}%`;
	'--background-opacity-multi'?: `${number}`;
} = $props();

export function open() {
	opened = true;
}
export function close() {
	opened = false;
}
export function toggle() {
	opened = !opened;
}

export function swipeend(e: CustomEvent<SwipeEndEventDetail>) {
	const { passThreshold } = e.detail;
	if (passThreshold) {
		close();
	}
}

export function setScrimOpacityMulti(multi: number) {
	document.documentElement.style.setProperty('--background-opacity-multi', String(multi));
}

beforeNavigate(async () => {
	opened = false;
});
const onkeypress = (event: KeyboardEvent) => {
	if (!opened) {
		return;
	}
	if (
		event.key === 'Enter' ||
		event.key === ' ' ||
		event.key === 'Spacebar' ||
		event.code === 'Escape' ||
		event.code === 'Enter' ||
		event.code === 'Space'
	) {
		opened = !opened;
	}
};
</script>

<svelte:window onkeypress={onkeypress} />

{#if opened}
	<div id="modal" transition:fade={{ duration: 196 }}>
		<button onclick={close} aria-label="modal-backdrop"></button>
		{@render children?.()}
	</div>
{/if}

<style lang="postcss">
#modal {
	position: absolute;
	inset: 0;
	isolation: isolate;
	z-index: 1000;
	&::after {
		content: '';
		position: absolute;
		pointer-events: none;
		inset: 0;
		z-index: -1;
		opacity: var(--background-opacity-multi, 1);
		backdrop-filter: var(--backdrop-filter, none);
		background:
			linear-gradient(
				color-mix(in srgb, transparent, rgb(32, 25, 47) var(--background-opacity, 80%)),
				color-mix(in srgb, transparent, rgb(7, 4, 11) var(--background-opacity, 95%))
			),
			url('/filter-noise.svg');
	}
}
button {
	-webkit-tap-highlight-color: transparent;
	border: none;
	border-radius: 0;
	position: absolute;
	inset: 0;
	padding: 0;
	opacity: 0;
	box-shadow: none;
	z-index: -1000;
}
</style>
