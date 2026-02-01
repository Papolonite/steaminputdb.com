<script lang="ts">
import Searchbar from '$lib/components/search/searchbar.svelte';
</script>

<main>
	<div>
		<div>
			<h1>
				<strong>Steam API backed</strong> community driven database of SteamInput configurations
			</h1>
			<p>SteamInputDB uses the same APIs as Steam itself</p>
			<p>That meansEvery configuration on Steam is available here!</p>
		</div>
		<div class="bar-container">
			<Searchbar />
		</div>
	</div>
</main>

<style lang="postcss">
main {
	padding: 1em;
	display: grid;
	place-items: center;
	& > div {
		display: grid;
		place-items: center;
	}
}

h1 strong {
	color: var(--highlight-color);
}
p {
	font-size: 1.2em;
	font-weight: bold;
}

.bar-container {
	padding: 2em;
	width: 100%;
	max-width: 36em;
}
</style>
