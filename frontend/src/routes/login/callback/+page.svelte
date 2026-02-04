<script lang="ts">
import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { page } from '$app/state';
import Spinner from '$lib/components/Spinner.svelte';
import Icon from '@iconify/svelte';
import { onMount } from 'svelte';
import { fade } from 'svelte/transition';

const searchParams = page.url.searchParams;
let loginErr = $state<Error | undefined>();

onMount(async () => {
	try {
		const resp = await fetch('?/validateLogin', {
			method: 'POST',
			body: JSON.stringify(
				searchParams.entries().reduce(
					(acc, [key, value]) => {
						if (key.startsWith('openid.')) {
							acc[key.split('openid.')[1]!] = value;
						}
						return acc;
					},
					{} as Record<string, string>
				)
			),
			headers: {
				'x-sveltekit-action': 'true'
			}
		});
		if (!resp.ok) {
			loginErr = await resp.json();
			return;
		}
		const data = await resp.json();
		if (!data) {
			loginErr = new Error('No data received from login validation');
			return;
		}
		if (data.type === 'failure') {
			loginErr = data;
			return;
		}
		goto(resolve('/'), {
			invalidateAll: true
		});
	} catch (e) {
		loginErr = e as Error;
	}
});
</script>

<main>
	{#if loginErr}
		<p transition:fade>
			Ouh no, Login Failed! <br />
			Please try again later.
			{JSON.stringify(loginErr)}
		</p>
	{:else}
		<div transition:fade class="spinner">
			<Spinner size="15em" thickness="0.3em" />
			<Icon icon="mdi:steam" style="height: 12em; width: 12em;" />
		</div>
	{/if}

	<!-- {#await waitLogin()}
		<div transition:fade class="spinner">
			<Spinner size="15em" thickness="0.3em" />
			<Icon icon="mdi:steam" style="height: 12em; width: 12em;" />
		</div>
	{:catch error}
		<p transition:fade>
			Ouh no, Login Failed! <br />
			Please try again later.
			{error?.message}
		</p>
	{/await} -->
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
