<script lang="ts">
import { onNavigate } from '$app/navigation';
import favicon from '$lib/assets/favicon.svg?url';
import Footer from '$lib/components/Footer.svelte';
import Header from '$lib/components/header/Header.svelte';

import { page } from '$app/state';
import 'unfonts.css';
import { links } from 'unplugin-fonts/head';
import '../css/main.pcss';
const { children } = $props();

onNavigate((navigation) => {
	if (!document.startViewTransition) {
		return;
	}

	// prevent view transition for same-page navigations,
	// there should not be a fucking transition if nothing changes... 🙄
	if (navigation.from?.url.pathname === navigation.to?.url.pathname) {
		return;
	}

	return new Promise((resolve) => {
		document.startViewTransition(async () => {
			resolve();
			await navigation.complete;
		});
	});
});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	{#each links as link (link?.attrs?.href)}
		{#if link?.attrs?.onload}
			<link
				{...link?.attrs || {}}
				onload={function () {
					this.rel = 'stylesheet';
				}} />
		{:else}
			<link {...link?.attrs || {}} />
		{/if}
	{/each}
	<link rel="canonical" href={page.url.toString()} />
	<meta property="og:url" content={page.url.toString()} />
	<meta property="og:site_name" content="SteamInputDB" />
	<meta name="robots" content="index, follow" />
</svelte:head>

<Header />
{@render children()}
<Footer />

<style lang="postcss">
:global(body) {
	display: grid;
	grid-template-rows: auto 1fr auto;
	min-height: 100dvh;
	max-width: 100dvw;
}

:global(main) {
	grid-row: 2 / span 1;
	grid-column: 1 / span 1;
}

:global(footer) {
	grid-row: 3 / span 1;
	grid-column: 1 / span 1;
}
</style>
