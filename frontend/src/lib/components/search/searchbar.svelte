<script lang="ts">
import IcOutlineSearch from '~icons/ic/outline-search';
let {
	value = $bindable()
}: {
	value?: unknown;
} = $props();
</script>

<label for="config-search">
	<input
		id="config-search"
		name="config-search"
		type="text"
		bind:value={value}
		placeholder="Search configurations or games..." />
	<IcOutlineSearch />
</label>

<style lang="postcss">
label {
	--font-size: 1.4em;

	position: relative;
	display: grid;
	grid-template-columns: auto min-content;
	place-items: center;

	border: 1px solid color-mix(in srgb, var(--text-color), transparent 90%);
	outline: 1px solid transparent;
	color: var(--text-color);
	border-radius: 100em;
	box-shadow: 0 1px 4px 0 rgb(0 0 0 / 0.25);
	transition: all 0.2s ease-in-out;
	&:is(:hover, :focus, :focus-within) {
		outline: 2px solid color-mix(in srgb, var(--color-primary), transparent 25%);
		& > :global(svg) {
			opacity: 0.75;
		}
	}
	background: var(--background-neutral-alpha);

	&::before {
		content: '';
		position: absolute;
		inset: 0;
		z-index: -1;
		background: var(--card-color);
		border-radius: 100em;
	}

	& > :global(svg) {
		width: var(--font-size);
		height: var(--font-size);
		margin-right: 1em;
		opacity: 0.5;
		transition: opacity var(--transition-duration) ease-in-out;
	}
}
input[type='text'] {
	font-size: var(--font-size);
	width: 100%;
	border: none;
	box-shadow: none;
	padding-right: 0.5em;
	padding-left: 1em;
	&:is(:hover, :focus) {
		outline: 1px solid transparent;
	}
}
</style>
