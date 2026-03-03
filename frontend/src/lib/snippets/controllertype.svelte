<script lang="ts" module>
export { controllertype };
</script>

<script lang="ts">
import type { components } from '$lib/api/openapi';
import SC2Icon from '$lib/assets/SC2_Alt.svg?component';
import Icon from '@iconify/svelte';
</script>

{#snippet controllertype({ item }: { item: components['schemas']['ConfigItem' | 'ConfigDetailResponse'] })}
	{#if item.controller_type === 'controller_neptune'}
		<Icon icon="simple-icons:steamdeck" width="1.2em" />
	{:else if item.controller_type === 'controller_triton'}
		<SC2Icon width="1.2em" />
	{:else if item.controller_type === 'controller_steamcontroller_gordon'}
		<SC2Icon width="1.2em" />
	{:else if item.controller_type === 'controller_ps5'}
		<Icon icon="simple-icons:playstation5" width="1.2em" />
	{:else if item.controller_type === 'controller_ps5_edge'}
		<Icon icon="simple-icons:playstation5" width="1.2em" />
	{:else if item.controller_type === 'controller_ps4'}
		<Icon icon="iconoir:playstation-gamepad" width="1.2em" />
	{:else if item.controller_type === 'controller_xbox360'}
		<Icon icon="fluent:xbox-controller-24-regular" width="1.2em" />
	{:else if item.controller_type === 'controller_xboxone'}
		<Icon icon="fluent:xbox-controller-24-filled" width="1.2em" />
	{:else if item.controller_type === 'controller_switch_pro'}
		<Icon icon="mdi:controller" width="1.2em" />
	{:else if item.controller_type === 'controller_8bitdo'}
		<Icon icon="mdi:controller" width="1.2em" />
	{:else if item.controller_type === 'controller_legion_go_s'}
		<Icon icon="game-icons:spartan-helmet" width="1.2em" />
	{:else if item.controller_type === 'controller_hori_steam'}
		<Icon icon="mdi:controller" width="1.2em" />
	{:else if item.controller_type === 'controller_rog_ally'}
		<Icon icon="simple-icons:republicofgamers" width="1.2em" />
	{:else if item.controller_type === 'controller_mobile_touch'}
		<Icon icon="mdi:cellphone" width="1.2em" />
	{:else if item.controller_type === 'controller_android'}
		<Icon icon="mdi:android" width="1.2em" />
	{:else}
		<Icon icon="mdi:gamepad" height="1.2em" />
	{/if}

	{item.controller_type_nice || item.controller_type || 'Generic Controller'}
{/snippet}
