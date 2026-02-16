<script lang="ts">
import Content from './content.svx';
</script>

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
	gap: 0.5em;
	height: fit-content;
	:global(h1, h2, h3, h4, h5, h6) {
		margin-top: 0.5em;
		margin-bottom: 0.2em;
	}
}
</style>
