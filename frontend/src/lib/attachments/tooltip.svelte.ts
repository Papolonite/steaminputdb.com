export const attachTooltip = (txtOrContainer?: string | Element) => (element: Element) => {
    let tooltip: HTMLElement | null = null;

    const showTooltip = () => {
        if (tooltip) {
            return;
        }

        tooltip = document.createElement('div');
        tooltip.style.position = 'fixed';
        tooltip.className = 'card';
        if (typeof txtOrContainer === 'string' || !txtOrContainer) {
            tooltip.textContent = txtOrContainer;
        } else {
            tooltip.appendChild(txtOrContainer.cloneNode(true));
        }
        document.body.appendChild(tooltip);

        const rect = element.getBoundingClientRect();
        const tooltipRect = tooltip.getBoundingClientRect();
        const top = rect.top - tooltipRect.height - 5;
        const left = rect.left + (rect.width - tooltipRect.width) / 2;
        tooltip.style.top = `${top}px`;
        tooltip.style.left = `calc(${left}px + 300px)`;
    };

    const hideTooltip = () => {
        if (tooltip) {
            document.body.removeChild(tooltip);
            tooltip = null;
        }
    };

    element.addEventListener('pointerenter', showTooltip);
    element.addEventListener('pointerleave', hideTooltip);

    return {
        update(newTxtOrContainer: string | Element) {
            txtOrContainer = newTxtOrContainer;
            if (tooltip) {
                if (typeof txtOrContainer === 'string') {
                    tooltip.textContent = txtOrContainer;
                } else {
                    tooltip.innerHTML = '';
                    tooltip.appendChild(txtOrContainer.cloneNode(true));
                }
            }
        },
        destroy() {
            element.removeEventListener('pointerenter', showTooltip);
            element.removeEventListener('pointerleave', hideTooltip);
            hideTooltip();
        }
    };

};
