<script lang="ts">
// LETZ M;E HAVE LINEBREAKS STUPID FUCK
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { page } from '$app/state';
import { client as apiclient } from '$lib/api/client';
import Spinner from '$lib/components/Spinner.svelte';
import Icon from '@iconify/svelte';
import { fade } from 'svelte/transition';

let openIdParams: Record<string, string> | undefined = $state();
if (browser) {
	openIdParams = page.url.searchParams.entries().reduce(
		(acc, [k, v]) => {
			if (k.startsWith('openid.')) {
				acc[k] = v;
				return acc;
			} else {
				return acc;
			}
		},
		{} as Record<string, string>
	);
}

const tryLogin = async () => {
	if (!openIdParams) {
		return;
	}
	const r = await apiclient.POST('/v1/steam/login', {
		body: Object.entries(openIdParams).reduce((acc, [k, v]) => {
			const key = k.split('.')?.pop();
			if (!key) {
				return acc;
			}
			acc[key] = v;
			return acc;
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
		}, {} as any)
	});
	if (r.error) {
		throw r.error;
	}
	if (r.data) {
		if (!r.data.token || !r.data.steam_id) {
			throw new Error('Invalid login response from server');
		}
		const mid = r.data.token.split('.')?.[1];
		if (!mid) {
			throw new Error('Invalid JWT token received');
		}
		const decoded = atob(mid);
		console.log('Decoded JWT:', JSON.parse(decoded));
		const expiresIn = JSON.parse(decoded).exp - Math.floor(Date.now() / 1000);
		console.log(`Token expires in ${expiresIn} seconds`);
		document.cookie = `token=${r.data.token};path=/;max-age=${expiresIn}`;
		document.cookie = `steamid=${r.data.steam_id};path=/;max-age=${expiresIn}`;
	}
	return r;
};
</script>

<main>
	{#if !openIdParams}
		<div transition:fade class="spinner">
			<Spinner size="15em" thickness="0.3em" />
			<Icon icon="mdi:steam" style="height: 12em; width: 12em;" />
		</div>
	{:else}
		{#await tryLogin()}
			<div transition:fade class="spinner">
				<Spinner size="15em" thickness="0.3em" />
				<Icon icon="mdi:steam" style="height: 12em; width: 12em;" />
			</div>
		{:then success}
			{success && ''}
			{goto(resolve('/'))}
		{:catch error}
			{error && ''}
			<!-- {goto(resolve('/error'))} -->
			Something went wrong during login. Please try again.
		{/await}
	{/if}
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
