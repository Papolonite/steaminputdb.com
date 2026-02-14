<script lang="ts">
import SC2 from '$lib/assets/SC2_Googley.svg.svelte';
import SearchForm from '$lib/components/search/searchform.svelte';
import { onMount } from 'svelte';

import Spinner from '$lib/components/Spinner.svelte';
import { fade } from 'svelte/transition';
import type { ActionData, PageData } from './$types';

let { data, form: formdata }: { data: PageData; form: ActionData } = $props();

$inspect(data, formdata);

let form = $state<HTMLFormElement>()!;
let formDisabled = $state(false);
let loading = $state(false);

let eyes = $state<{
	left: HTMLElement;
	right: HTMLElement;
}>({
	left: undefined!,
	right: undefined!
})!;
onMount(() => {
	const group = document.querySelector('#sc2>*>svg>g>:last-child')!;
	eyes.left = group.children[0] as HTMLElement;
	eyes.right = group.children[1] as HTMLElement;
});
</script>

<svelte:window
	onmousemove={(e) => {
		if (!eyes || !eyes.left || !eyes.right) {
			return;
		}
		Object.values(eyes).forEach((eye) => {
			const rect = eye.getBoundingClientRect();
			const eyeX = rect.left + rect.width / 2;
			const eyeY = rect.top + rect.height / 2;
			const deltaX = e.clientX - eyeX;
			const deltaY = e.clientY - eyeY;
			const angle = Math.atan2(deltaY, deltaX);
			const distance = Math.min(16, Math.hypot(deltaX, deltaY) / 2);
			const translateX = Math.cos(angle) * distance;
			const translateY = Math.sin(angle) * distance;
			eye.style.transform = `translate(${translateX}px, ${translateY}px)`;
		});
	}} />

<main>
	<search>
		<SearchForm
			bind:form={form}
			disabled={formDisabled}
			method="POST"
			action="?/search"
			enhanceParams={() => {
				loading = true;

				return async ({ update }) => {
					formDisabled = true;
					loading = true;
					try {
						await update({ reset: false });
					} finally {
						formDisabled = false;
						loading = false;
					}
				};
			}} />
		{#if loading}
			<div id="sc2" transition:fade>
				<Spinner size="16em" />
			</div>
		{:else if !formdata}
			<div id="sc2" transition:fade>
				<SC2
					height="100%"
					--eyes-color="black"
					--eyes-white-color="var(--text-color-dark)"
					--eyes-border-color="light-dark(var(--text-color-light), transparent)" />
			</div>
		{:else if formdata.results?.items?.length == 0}
			<div id="sc2" transition:fade>
				<span>No results found</span>
				<SC2
					height="100%"
					--eyes-color="black"
					--eyes-white-color="var(--text-color-dark)"
					--eyes-border-color="light-dark(var(--text-color-light), transparent)" />
			</div>
		{:else}
			<div class="results" transition:fade>
				<!-- render entries -->
			</div>
		{/if}
	</search>
</main>

<style lang="postcss">
main {
	display: grid;
	place-items: center;
}

search {
	display: grid;
	place-self: center;
	gap: 1em;
	place-items: center;
	min-width: 50%;
	--max-width: 1440px;
	max-width: min(100%, var(--max-width));
	isolation: isolate;

	transition: all var(--transition-duration) var(--default-ease) allow-discrete;

	& > :first-child {
		width: 100%;
	}
	margin-top: 1em;
	height: 100%;
}

.results {
	display: grid;
}

#sc2 {
	filter: drop-shadow(0px 0.25em 0.2em var(--shadow-color));
	height: 100%;
	width: 100%;
	display: grid;
	place-items: center;
	padding: 0.5em;

	:global(svg) {
		max-height: 18em;
		margin-bottom: auto;
	}

	@media (any-pointer: coarse) {
		:global(ellipse) {
			transition: transform calc(var(--transition-duration) * 2) var(--default-ease);
		}
	}
}
</style>
