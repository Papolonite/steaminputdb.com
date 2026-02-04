<script lang="ts">
import { onNavigate } from '$app/navigation';
import favicon from '$lib/assets/favicon.svg';
import Footer from '$lib/components/Footer.svelte';
import Header from '$lib/components/Header.svelte';
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
}
</style>
