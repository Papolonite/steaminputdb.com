<script lang="ts">
import Content from './content.svx';
</script>

<svelte:head>
	<title>SteamInputDB - Imprint</title>
	<meta property="og:title" content="SteamInputDB - Imprint" />
	<meta name="description" content="Imprint of SteamInputDB" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="SteamInputDB - Imprint" />
</svelte:head>

<main>
	<div>
		<Content />
	</div>
</main>

<style lang="postcss">
main {
	padding: 2em;
	display: grid;
	gap: 0.75em;
	height: 100%;
	min-width: 50%;
	--max-width: 1440px;
	max-width: min(100%, var(--max-width));
	place-self: center;
}
div {
	display: grid;
	text-align: justify;
	height: fit-content;
	:global(h1),
	:global(h2),
	:global(h3),
	:global(h4),
	:global(h5),
	:global(h6) {
		margin-top: 1em;
		margin-bottom: 0.5em;
	}
}
</style>
