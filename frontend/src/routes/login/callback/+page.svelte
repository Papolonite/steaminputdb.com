<script lang="ts">
import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { page } from '$app/state';
import Spinner from '$lib/components/Spinner.svelte';
import Icon from '@iconify/svelte';
import type { ActionFailure } from '@sveltejs/kit';
import { fade } from 'svelte/transition';

const waitLogin = async () => {
	const cookieOrFail: string | ActionFailure = await page.data.loginPromise;

	if (typeof cookieOrFail !== 'string') {
		throw cookieOrFail;
	}

	document.cookie = cookieOrFail;
	goto(resolve('/'), {
		invalidateAll: true
	});
};
</script>

<main>
	{#await waitLogin()}
		<div transition:fade class="spinner">
			<Spinner size="15em" thickness="0.3em" />
			<Icon icon="mdi:steam" style="height: 12em; width: 12em;" />
		</div>
	{:catch error}
		<p transition:fade>
			Ouh no, Login Failed! <br />
			Please try again later.
			{error?.message}
			<!-- TODO: create proper error page -->
		</p>
	{/await}
</main>

<style lang="postcss">
main {
	display: grid;
	place-items: center;
	& > :global(*) {
		grid-area: 1 / 1;
	}
	& .spinner {
		display: grid;
		place-items: center;
		width: 100%;
		height: 50%;
		& > :global(*) {
			grid-area: 1 / 1;
		}
	}
}
</style>
