<script lang="ts">
import favicon from '$lib/assets/favicon.svg';
import Footer from '$lib/components/Footer.svelte';
import Header from '$lib/components/Header.svelte';
import { applyUserThemePreferences } from '$lib/components/theme/applytheme.svelte';
import { onMount } from 'svelte';
import '../css/main.pcss';

const applyTransitionDuration = applyUserThemePreferences();
onMount(() => {
	applyTransitionDuration?.();
});

const { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Header />
{@render children()}
<div></div>
<Footer />

<style lang="postcss">
:global(body) {
	display: grid;
	grid-template-rows: auto auto 1fr auto;
	min-height: 100dvh;
}
</style>
