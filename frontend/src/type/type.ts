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
export type AnswerDetailResponse = {
	id: string;
	questionId: string;
	participantId: string;
	content: string;
	attachments: AttachmentResponse[];
	createdAt: string;
};
export type QuestionDetailResponse = {
	id: string;
	participantId: string;
	content: string;
	attachments: AttachmentResponse[];
	createdAt: string;
};
export type ChatDetailResponse = {
	id: string;
	title?: string | null;
	participants: ParticipantResponse[];
	questions: QuestionDetailResponse[];
	answers: AnswerDetailResponse[];
	startedAt: string;
	lastActiveAt: string;
};
