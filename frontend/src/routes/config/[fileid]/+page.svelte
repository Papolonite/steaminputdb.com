<script lang="ts">
import { page } from '$app/state';
import type { components } from '$lib/api/openapi';
import { sectionHead } from './sectionHead.svelte';
import { sectionInfo } from './sectionInfo.svelte';

const fileInfo: components['schemas']['ConfigDetailResponse'] = $derived(page.data.fileInfo);
const appInfo: components['schemas']['AppInfo'] = $derived(page.data.appInfo);

$inspect(page);
</script>

<main style={appInfo?.background ? `--bg: url('${appInfo?.background}')` : ''}>
	<div>
		{@render sectionHead({ fileInfo, appInfo })}
		{@render sectionInfo({ fileInfo, appInfo })}
	</div>
</main>

<style lang="postcss">
main {
	position: relative;
	isolation: isolate;
	display: grid;
	padding: 1em 0;

	place-items: center;
	grid-template-rows: min-content;
	grid-template-columns: minmax(min(100%, auto), 50%);
	width: 100%;

	&::before {
		content: '';
		position: absolute;
		inset: 0;
		background:
			linear-gradient(
				0deg,
				var(--background-color) 25%,
				color-mix(in srgb, var(--background-color), transparent 100%) 80%
			),
			var(--bg, transparent) top/cover no-repeat;
		z-index: -2;
	}
	&::after {
		content: '';
		position: absolute;
		inset: 0;
		z-index: -1;
		backdrop-filter: blur(12px);
	}
}
div {
	display: grid;
	place-self: center;
	gap: 1em;
	place-items: center;
	min-width: 50%;
	--max-width: 1440px;
	max-width: min(100%, var(--max-width));
	isolation: isolate;
	/* container: main / inline-size;*/
	:global(> :first-child) {
		z-index: -1;
		width: 100%;
	}
}
</style>
