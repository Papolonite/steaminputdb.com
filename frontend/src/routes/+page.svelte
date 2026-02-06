<script lang="ts">
import SC2 from '$lib/assets/SC2_Googley.svg?component';
import { onMount } from 'svelte';
let eyes = $state<{
	left: HTMLElement;
	right: HTMLElement;
}>({
	left: undefined!,
	right: undefined!
})!;

onMount(() => {
	const group = document.querySelector('#sc2>*>svg>g>:last-child')!;
	eyes.left = group.children[0] as HTMLElement;
	eyes.right = group.children[1] as HTMLElement;
});
</script>

<svelte:window
	onmousemove={(e) => {
		if (!eyes || !eyes.left || !eyes.right) {
			return;
		}
		Object.values(eyes).forEach((eye) => {
			const rect = eye.getBoundingClientRect();
			const eyeX = rect.left + rect.width / 2;
			const eyeY = rect.top + rect.height / 2;
			const deltaX = e.clientX - eyeX;
			const deltaY = e.clientY - eyeY;
			const angle = Math.atan2(deltaY, deltaX);
			const distance = Math.min(16, Math.hypot(deltaX, deltaY) / 2);
			const translateX = Math.cos(angle) * distance;
			const translateY = Math.sin(angle) * distance;
			eye.style.transform = `translate(${translateX}px, ${translateY}px)`;
		});
	}} />

<main>
	<div>
		<div>
			<div id="sc2">
				<SC2
					height="100%"
					--eyes-color="black"
					--eyes-white-color="var(--text-color-dark)"
					--eyes-border-color="light-dark(var(--text-color-light), transparent)" />
			</div>
			<div>
				<h1>
					Community driven database of <span>SteamInput</span> configurations
					<strong>utilizing Steam API</strong>
				</h1>
				<p>SteamInputDB uses the same APIs as Steam itself</p>
				<p>That means <strong>every</strong> configuration on Steam is also available here!</p>
			</div>
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

main > div > :first-child {
	display: grid;
	place-items: center;
}

h1 {
	& strong {
		color: var(--color-primary);
	}
	& span {
		color: var(--highlight-color);
	}
}
p {
	font-size: 1.2em;
	font-weight: bold;
	& strong {
		color: var(--highlight-color);
	}
}

#sc2 {
	filter: drop-shadow(0px 0.25em 0.2em var(--shadow-color));
	max-height: 16em;
	height: 100%;
	width: 100%;
	display: grid;
	place-items: center;
	padding: 0.5em;
	@media (any-pointer: coarse) {
		:global(ellipse) {
			transition: transform calc(var(--transition-duration) * 2) var(--default-ease);
		}
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
