<script lang="ts">
import favicon from '$lib/assets/favicon.svg';
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

{@render children()}
