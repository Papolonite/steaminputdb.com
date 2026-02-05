<script lang="ts">
import { resolve } from '$app/paths';
</script>

<footer>
	<div class="disclaimer">
		<div>
			<span>SteamInputDB is not affiliated with Steam, Valve Software or any of their partners.</span>
			<span>
				Steam and the Steam logo are trademarks of Valve Corporation. All other trademarks are
				property of their respective owners.
			</span>
		</div>
		<a href={resolve('/imprint')}>Imprint</a>
	</div>
</footer>

<style lang="postcss">
footer {
	padding: 1em;
	background: var(--card-background-noise);
	box-shadow: 0 0px 4px var(--shadow-color);
}

.disclaimer {
	display: grid;
	width: 100%;
	grid-template-columns: auto min-content;
	align-items: end;
	& > :first-child {
		display: grid;
		grid-template-columns: 1fr;
	}
	font-size: 0.8em;
	color: var(--text-muted);
	a {
		color: inherit;
	}
}
</style>
