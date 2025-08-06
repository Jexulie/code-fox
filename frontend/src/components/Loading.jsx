import {Spinner} from "@heroui/react";

export default function Loading({size='md', color='primary', variant='default', label=undefined}) {
    return (
        <Spinner size={size} color={color} label={label} variant={variant} />
    )
}