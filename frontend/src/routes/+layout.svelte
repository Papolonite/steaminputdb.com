<script lang="ts">
import favicon from '$lib/assets/favicon.svg';
import '../css/main.pcss';

const { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children()}
