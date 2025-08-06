import {Modal, ModalBody, ModalContent, ModalFooter, ModalHeader} from "@heroui/modal";
import {Button} from "@heroui/react";


export default function DialogModal({isOpen, onOpenChange, content, onAccept}) {
    const handleAgree = async () => {
        await onAccept();
    }

    return (
        <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
            <ModalContent>
                {(onClose) => (
                    <>
                        <ModalHeader className="flex flex-col gap-1 text-gray-500">Are you sure ?</ModalHeader>
                        <ModalBody className='w-full'>
                            {content.map((item, index) => (
                                <p key={index} className='text-gray-500 text-center'>{item}</p>
                            ))}

                        </ModalBody>
                        <ModalFooter className='w-full'>
                            <Button color="default" onPress={async () => {
                                await handleAgree();
                                onClose();
                            }} fullWidth>
                                Yes
                            </Button>
                            <Button color="danger" onPress={onClose} fullWidth>
                                Close
                            </Button>
                        </ModalFooter>
                    </>
                )}
            </ModalContent>
        </Modal>
    )
}