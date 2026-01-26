/** @type {import('postcss-load-config').Config} */
import cssnano from 'cssnano'
import presetEnv from 'postcss-preset-env'
const config = {
    plugins: [
        presetEnv({
            autoprefixer: true,
            stage: 3,
            features: {
                'nesting-rules': true,
            }
        }),
        cssnano({
            preset: 'default',
        })
    ]
}

export default config