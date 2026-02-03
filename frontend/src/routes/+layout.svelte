<script lang="ts">
import favicon from '$lib/assets/favicon.svg';
import Footer from '$lib/components/Footer.svelte';
import Header from '$lib/components/Header.svelte';
import '../css/main.pcss';

const { children } = $props();
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
