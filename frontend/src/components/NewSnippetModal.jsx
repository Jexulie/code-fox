import {useState} from "react";
import {Modal, ModalBody, ModalContent, ModalFooter, ModalHeader} from "@heroui/modal";
import {Button, Form, Input} from "@heroui/react";


export default function NewSnippetModal({isOpen, onOpenChange, addPassword}) {
    const [errors, setErrors] = useState({});

    const onSubmit = async (e, onClose) => {

        const data = Object.fromEntries(new FormData(e.currentTarget));

        addPassword(data.service, data.username, data.password);

        setErrors({});
        onClose()
    }

    return (
        <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
            <ModalContent>
                {(onClose) => (
                    <>
                        <Form
                            id='new-password'
                            validationErrors={errors}
                            onSubmit={async (e) => {
                                e.preventDefault();
                                await onSubmit(e, onClose);
                            }}
                            onReset={() => false}>
                            <ModalHeader className="flex flex-col gap-1 text-gray-500">New Password</ModalHeader>
                            <ModalBody className='w-full'>
                                <Input
                                    isRequired
                                    type="service"
                                    name="service"
                                    label="Service"
                                    errorMessage={({validationDetails}) => {
                                        if (validationDetails.valueMissing) {
                                            return "Service is missing";
                                        }

                                        return errors.service;
                                    }}
                                    fullWidth/>

                                <Input
                                    type="username"
                                    name="username"
                                    label="Username"
                                    errorMessage={({validationDetails}) => {
                                        if (validationDetails.valueMissing) {
                                            return "Username is missing";
                                        }

                                        return errors.username;
                                    }}
                                    fullWidth/>

                                <Input
                                    isRequired
                                    type="password"
                                    name="password"
                                    label="Password"
                                    errorMessage={({validationDetails}) => {
                                        if (validationDetails.valueMissing) {
                                            return "Password is missing";
                                        }

                                        return errors.password;
                                    }}
                                    fullWidth/>

                            </ModalBody>
                            <ModalFooter className='w-full'>
                                <Button color="default" type="submit" fullWidth>
                                    Save
                                </Button>
                                <Button color="danger" onPress={onClose} fullWidth>
                                    Close
                                </Button>
                            </ModalFooter>
                        </Form>
                    </>
                )}
            </ModalContent>
        </Modal>
    )
}