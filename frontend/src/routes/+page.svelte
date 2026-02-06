<script lang="ts">
</script>

<main>
	<div>
		<div>
			<h1>
				<strong>Steam API backed</strong> community driven database of SteamInput configurations
			</h1>
			<p>SteamInputDB uses the same APIs as Steam itself</p>
			<p>That means <strong>every</strong> configuration on Steam is also available here!</p>
		</div>
		<div class="wip">
			<span>🚧 Work in Progress 🚧</span>
			<span>More coming Soon™</span>
		</div>
	</div>
</main>

<style lang="postcss">
main {
	padding: 2em;
	display: grid;
	place-items: center;
	position: relative;
	& > div {
		display: grid;
		place-items: center;
	}
}

h1 strong {
	color: var(--color-primary);
}
p {
	font-size: 1.2em;
	font-weight: bold;
	& strong {
		color: var(--highlight-color);
	}
}

.wip {
	display: grid;
	place-items: center;
	margin-top: 2em;
	gap: 1em;
	& :first-child {
		font-size: 1.8em;
		@media screen and (max-width: 400px) {
			font-size: 1.5em;
		}
	}
	& :last-child {
		color: var(--highlight-color);
		font-size: 2em;
		font-weight: bold;
	}
}
</style>
