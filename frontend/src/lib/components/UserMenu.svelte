<script lang="ts">
import { page } from '$app/state';

const userInfo = $state(page.data.userInfo);
</script>

<a href="TODO">
	{#if userInfo?.avatar}
		<enhanced:img src={userInfo.avatar} alt="User Avatar" height="100%"></enhanced:img>
	{:else}
		<span>{userInfo?.personaname?.split(/(\.|\[|\]|\s)/g)?.pop()?.[0] || '¯\\_(ツ)_/¯'}</span>
	{/if}
</a>

<style lang="postcss">
a {
	display: grid;
	place-items: center;
	min-width: 2em;
	height: 100%;
	aspect-ratio: 1 /1;

	box-shadow: 0 0 1em -0.5em var(--color-primary);
	border: 1px solid color-mix(in srgb, var(--color-primary), transparent 60%);
	cursor: pointer;
	transition: all var(--transition-duration) var(--default-ease);
	padding: 0;
	background: var(--color-primary);

	&:hover,
	&:focus,
	&:focus-within {
		box-shadow: 0 0 1em -0.3em var(--color-primary);
		transform: scale(1.05);
		outline: 2px solid color-mix(in srgb, var(--color-primary), transparent 60%);
	}
	&::before {
		display: none;
	}
	&::after {
		display: none;
	}
	& > span {
		font-size: 1.2em;
		padding: 0 25%;
		font-weight: bold;
		color: var(--text-color-dark);
	}
}
</style>
