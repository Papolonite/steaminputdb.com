

export const intersectionObserver = (callback = () => { }, threshold = 0.5) => (element: HTMLElement) => {
    const observer = new IntersectionObserver(
        (entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) {
                    callback();
                }
            });
        },
        {
            threshold: threshold
        }
    );
    observer.observe(element);
    return () => {
        observer.disconnect();
    };
};

