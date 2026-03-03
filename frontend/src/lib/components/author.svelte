<script lang="ts">
import DefaultAuthorImg from '$lib/assets/author_pr.jpg?enhanced';
import type { Picture } from '@sveltejs/enhanced-img';
import { format } from 'date-fns';

const {
	author = {
		name: 'Peter Repukat',
		img: DefaultAuthorImg
	},
	date = undefined
}: {
	author?: {
		name: string;
		img?: Picture;
	};
	date?: Date;
} = $props();
</script>

<div>
	{#if author.img}
		<enhanced:img src={author.img} alt="A picture of {author.name}, the author of this post"
		></enhanced:img>
	{:else}
		<div></div>
	{/if}
	<div>
		<span>{author.name}</span>
		{#if date}
			<span class="date">{format(date, 'PPP')}</span>
		{/if}
	</div>
</div>

<style lang="postcss">
div {
	& :global(picture),
	& :global(img) {
		width: 2.4em;
		height: 2.4em;
		border-radius: 100dvw;
		object-fit: cover;
		aspect-ratio: 1 / 1;
	}
	display: flex;
	flex-flow: row nowrap;
	gap: 1em;
	align-items: center;
	& span {
		font-size: 0.8em;
		font-weight: 500;
	}
	& .date {
		opacity: 0.8;
		font-weight: normal;
	}
	& div {
		display: grid;
		gap: 0.1em;
	}
}
</style>
