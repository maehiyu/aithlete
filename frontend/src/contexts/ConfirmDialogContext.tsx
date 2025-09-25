import React, { createContext, useContext, useState, ReactNode } from 'react';
import { Fragment } from 'react';
import { Dialog, DialogPanel, Transition, TransitionChild } from '@headlessui/react';
import { 
	ExclamationTriangleIcon, 
	InformationCircleIcon, 
	XCircleIcon 
} from '@heroicons/react/24/outline';
type ConfirmDialogData = {
	title: string;
	message: string;
	confirmText?: string;
	cancelText?: string;
	type?: 'info' | 'warning' | 'error';
};

type ConfirmDialogContextType = {
	confirm: (data: ConfirmDialogData) => Promise<boolean>;
};

const ConfirmDialogContext = createContext<ConfirmDialogContextType | undefined>(undefined);

export const useConfirm = () => {
	const context = useContext(ConfirmDialogContext);
	if (!context) {
		throw new Error('useConfirm must be used within a ConfirmDialogProvider');
	}
	return context;
};

type ConfirmDialogProviderProps = {
	children: ReactNode;
};

export const ConfirmDialogProvider = ({ children }: ConfirmDialogProviderProps) => {
	const [dialogData, setDialogData] = useState<ConfirmDialogData | null>(null);
	const [resolvePromise, setResolvePromise] = useState<((value: boolean) => void) | null>(null);

	const confirm = (data: ConfirmDialogData): Promise<boolean> => {
		return new Promise((resolve) => {
			setDialogData(data);
			setResolvePromise(() => resolve);
		});
	};

	const handleConfirm = () => {
		if (resolvePromise) {
			resolvePromise(true);
		}
		closeDialog();
	};

	const handleCancel = () => {
		if (resolvePromise) {
			resolvePromise(false);
		}
		closeDialog();
	};

	const closeDialog = () => {
		setDialogData(null);
		setResolvePromise(null);
	};

	return (
		<ConfirmDialogContext.Provider value={{ confirm }}>
			{children}
			{dialogData && (
				<ConfirmDialog
					{...dialogData}
					onConfirm={handleConfirm}
					onCancel={handleCancel}
				/>
			)}
		</ConfirmDialogContext.Provider>
	);
};


type ConfirmDialogProps = ConfirmDialogData & {
	onConfirm: () => void;
	onCancel: () => void;
};

const ConfirmDialog = ({ 
	title, 
	message, 
	confirmText = "確認", 
	cancelText = "キャンセル",  
	type = "info",
	onConfirm, 
	onCancel 
}: ConfirmDialogProps) => {
	const getIcon = () => {
		switch (type) {
			case 'warning':
				return <ExclamationTriangleIcon className="w-6 h-6 text-gray-600" />;   
			case 'error':
				return <XCircleIcon className="w-6 h-6 text-gray-600" />;
			default:
				return <InformationCircleIcon className="w-6 h-6 text-gray-600" />;
		}
	};

	const getConfirmButtonStyle = () => {
		return "bg-gray-900 hover:bg-gray-800 focus:ring-gray-500";
	};

	return (
		<Transition appear show={true} as={Fragment}>
			<Dialog as="div" className="relative z-50" onClose={onCancel}>
				<TransitionChild
					as={Fragment}
					enter="ease-out duration-300"
					enterFrom="opacity-0"
					enterTo="opacity-100"
					leave="ease-in duration-200"
					leaveFrom="opacity-100"
					leaveTo="opacity-0"
				>
					<div className="fixed inset-0 bg-black bg-opacity-25" />
				</TransitionChild>

				<div className="fixed inset-0 overflow-y-auto">
					<div className="flex min-h-full items-center justify-center p-4 text-center">
						<TransitionChild
							as={Fragment}
							enter="ease-out duration-300"
							enterFrom="opacity-0 scale-95"
							enterTo="opacity-100 scale-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100 scale-100"
							leaveTo="opacity-0 scale-95"
						>
							<DialogPanel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
								<div className="flex items-start space-x-4">
									<div className="flex-shrink-0">
										{getIcon()}
									</div>
									<div className="flex-1">
										<h3 className="text-lg font-medium leading-6 text-gray-900 mb-2">
											{title}
										</h3>
										<p className="text-sm text-gray-500">
											{message}
										</p>
									</div>
								</div>

								<div className="mt-6 flex justify-end space-x-3">
									<button
										type="button"
										className="inline-flex justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
										onClick={onCancel}
									>
										{cancelText}
									</button>
									<button
										type="button"
										className={`inline-flex justify-center rounded-md border border-transparent px-4 py-2 text-sm font-medium text-white focus:outline-none focus:ring-2 focus:ring-offset-2 ${getConfirmButtonStyle()}`}
										onClick={onConfirm}
									>
										{confirmText}
									</button>
								</div>
							</DialogPanel>
						</TransitionChild>
					</div>
				</div>
			</Dialog>
		</Transition>
	);
};
