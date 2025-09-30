export type ChatCreateRequest = {
	title?: string | null;
	participantIds: string[];
	questions?: string[];
};
export type ChatUpdateRequest = {
	id: string;
	title?: string | null;
	opponentIds?: string[];
};
export type ParticipantCreateRequest = {
	name: string;
	email: string;
	role: string;
	sports: string[];
	iconUrl?: string | null;
};
export type ParticipantUpdateRequest = {
	name?: string;
	email?: string;
	role?: string;
	sports?: string[];
	iconUrl?: string | null;
};
export type ParticipantResponse = {
	id: string;
	name: string;
	email: string;
	role: string;
	iconUrl?: string | null;
	sports: string[];
};
export type OpponentResponse = {
	id: string;
	name: string;
	role: string;
	iconUrl: string;
};
export type ChatSummaryResponse = {
	id: string;
	title?: string | null;
	lastActiveAt: string;
	latestQa?: string | null;
	opponent: OpponentResponse;
};
export type PoseDataResponse = {
	keypoints: string;
	score: number;
};
export type AttachmentResponse = {
	type: string;
	url: string;
	pose?: PoseDataResponse | null;
};

export type ChatItem = {
	id: string;
	chatId: string;
	participantId: string;
	content: string;
	createdAt: string;
	attachments: AttachmentResponse[];
	type: "question" | "answer" | "ai_answer";
	questionId?: string | null;
	status?: MessageStatus;
	tempId?: string;
};

export type ChatItemRequest = {
	participantId: string;
	questionId?: string;
	content: string;
	type: "question" | "answer" | "ai_answer";
	tempId?: string;
}

export type MessageStatus = "sending" | "sent" | "failed";

export type WithStatus<T> = T & {
	status: MessageStatus;
	tempId?: string;
};
export type ChatDetailResponse = {
	id: string;
	title?: string | null;
	participants: ParticipantResponse[];
	timeline: ChatItem[];
	startedAt: string;
	lastActiveAt: string;
};

export type AppointmentCreateRequest = {
	chatId: string;
	title: string;
	description: string;
	scheduledAt: string;
	duration: number;
	participantIds: string[];
};

export type AppointmentUpdateRequest = {
	title?: string | null;
	description?: string | null;
	scheduledAt?: string;
	duration?: number;
	status?: "scheduled" | "completed" | "canceled";
};

export type AppointmentResponse = {
	id: string;
	chatId: string;
	title: string;
	description: string;
	scheduledAt: string;
	duration: number;
	status: "scheduled" | "completed" | "canceled";
	participants: AppointmentParticipantResponse[];
}

export type AppointmentParticipantResponse = {
	participantId: string;
	participationStatus: "needs-action" | "accepted" | "declined" | "tentative";
};
