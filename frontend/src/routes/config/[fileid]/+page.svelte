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
	{@render sectionHead({ fileInfo, appInfo })}
	{@render sectionInfo({ fileInfo, appInfo })}
</main>

<style lang="postcss">
main {
	position: relative;
	isolation: isolate;
	display: grid;
	justify-content: center;
	padding: 1em 0;
	min-width: 50%;
	max-width: 100%;
	gap: 1em;
	grid-template-rows: min-content min-content;

	&::before {
		content: '';
		position: absolute;
		inset: 0;
		background:
			linear-gradient(0deg, var(--background-color) 25%, transparent 80%),
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
</style>
