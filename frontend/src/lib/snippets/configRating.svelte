<script lang="ts" module>
export { configRating };
</script>

<script lang="ts">
import type { components } from '$lib/api/openapi';
import { tooltip } from '$lib/attachments/tooltip.svelte';
import Icon from '@iconify/svelte';
</script>

{#snippet configRating({ item }: { item: components['schemas']['ConfigItem' | 'ConfigDetailResponse'] })}
	{#if item.votes}
		{@const scoreColor =
			item.votes.up == item.votes.up && item.votes.up === 0
				? 'currentColor'
				: item.votes?.score &&
					`hsl(${
						(item.votes.score || 0) > 0.8
							? 125
							: (item.votes.score || 0) > 0.7
								? 80
								: (item.votes.score || 0) > 0.6
									? 60
									: (item.votes.score || 0) > 0.5
										? 30
										: 0
					}deg 100% 50%)`}
		<div
			class="rating"
			style="--rating-color: {scoreColor};"
			{@attach tooltip({
				snippet: tooltipContent,
				snippetInDefaultBackground: true,
				outDelay: 200,
				arrow: true,
				arrowFollowCursor: true
			})}>
			<div>
				<span>
					{#if (item.votes?.score || 0) > 0.8}
						😍
					{:else if (item.votes?.score || 0) > 0.7}
						🤩
					{:else if (item.votes?.score || 0) > 0.6}
						😎
					{:else if (item.votes?.score || 0) > 0.5}
						🙁
					{:else if (item.votes?.down || 0) > (item.votes?.up || 0)}
						😣
					{:else}
						🤔
					{/if}
				</span>
				<span>
					{item.votes.up == item.votes.down && item.votes.up === 0
						? '???'
						: (item.votes?.score ?? 0).toLocaleString(undefined, {
								style: 'percent',
								minimumFractionDigits: 0,
								maximumFractionDigits: 1
							})}
				</span>
			</div>
			<div>
				<span
					>{item.votes.up?.toLocaleString(undefined, {
						notation: 'compact',
						minimumFractionDigits: 0,
						maximumFractionDigits: 1
					})}
					<Icon icon="mdi:thumb-up" />
				</span><span>/</span><span>
					{item.votes.down?.toLocaleString(undefined, {
						notation: 'compact',
						minimumFractionDigits: 0,
						maximumFractionDigits: 1
					})}
					<Icon icon="mdi:thumb-down" /></span>
			</div>
			<i>(Rating from Steam)</i>
		</div>
	{/if}
{/snippet}

{#snippet tooltipContent()}
	<div style="display: grid; place-items: center;">
		<p style="white-space: nowrap; text-align: center;">The ranking system is provided by Steam</p>
		<p style="white-space: nowrap; text-align: center;">
			I do not know and can only guess on how it rates
		</p>
	</div>
{/snippet}

<style lang="postcss">
.rating {
	display: grid;
	place-items: center;
	& span {
		filter: drop-shadow(1px 1px 1px black);
		color: var(--rating-color);
		font-size: 1.1em;
		font-weight: 500;
	}
	& i {
		color: var(--text-color-dark);
		opacity: 0.8;
		font-size: 0.8em;
		filter: drop-shadow(1px 1px 1px black) drop-shadow(0px 0px 2px black);
	}
	& > :first-child {
		& > :first-child {
			font-size: 1.8em;
			filter: drop-shadow(1px 2px 3px black);
			transform: translate(0, 0.5em);
		}
		& > :last-child {
			font-size: 1.7em;
			filter: drop-shadow(1px 1px 1px black) drop-shadow(1px 1px 2px var(--shadow-color));
		}
	}
	& > :nth-child(2) {
		display: flex;
		gap: 0.5ch;
		& > :first-child {
			color: hsl(108, 100%, 50%);
		}
		& > :nth-child(2) {
			font-size: 1.1em;
			color: var(--text-color-dark);
		}
		& > :last-child {
			color: hsl(0, 100%, 60%);
		}
	}
	height: 100%;
}
</style>
