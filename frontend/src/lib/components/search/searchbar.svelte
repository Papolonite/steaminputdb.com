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
	box-shadow: 0 1px 4px 0 var(--shadow-color);
	transition: all var(--transition-duration) var(--default-ease);

	&:hover,
	&:focus-within {
		outline: 0.1em solid var(--color-primary);
		box-shadow: 0 0 1.3em -0.4em var(--color-primary);
		& :global(svg) {
			opacity: 1;
		}
	}

	background: var(--card-color);
	isolation: isolate;
	&::before {
		content: '';
		position: absolute;
		inset: 0;
		z-index: -1;
		background: var(--background-neutral-alpha);
		border-radius: 100em;
	}

	& > :global(svg) {
		width: var(--font-size);
		height: var(--font-size);
		margin-right: 1em;
		opacity: 0.5;
		transition: all var(--transition-duration) var(--default-ease);
	}
}
input[type='text'] {
	font-size: var(--font-size);
	width: 100%;
	border: none;
	box-shadow: none;
	padding-right: 0.5em;
	padding-left: 1em;

	outline: none;
	&:is(:hover, :focus) {
		outline: none;
	}
}
</style>
