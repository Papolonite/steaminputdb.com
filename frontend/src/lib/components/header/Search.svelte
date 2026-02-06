<script lang="ts">
import Searchbar from '../search/searchbar.svelte';
</script>

<div>
	<Searchbar --box-shadow="inset 0 0.4em 0.4em 0 var(--shadow-color-dark)" />
</div>

<style lang="postcss">
div {
	place-self: center;
	display: grid;
	grid-template-columns: minmax(16ch, 48ch);
}
</style>
