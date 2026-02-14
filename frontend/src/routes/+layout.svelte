<script lang="ts">
import { onNavigate } from '$app/navigation';
import favicon from '$lib/assets/favicon.svg?url';
import Footer from '$lib/components/Footer.svelte';
import Header from '$lib/components/header/Header.svelte';
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
</style>
