<script lang="ts">
let form = $state<HTMLFormElement>()!;
</script>

<main>
	<search>
		<form bind:this={form}>
			<label for="config-search">
				<input
					type="search"
					id="config-search"
					name="config-search"
					placeholder="SteamInput configuration..." />
				<button type="button">search</button>
			</label>
			<input type="range" min="0" max="100" value="50" />
		</form>
		<div class="results">
			<!-- render entries -->
		</div>
	</search>
</main>

<style lang="postcss">
main {
	display: grid;
	place-items: center;
}

search {
	display: grid;
}

.results {
	display: grid;
}
</style>
