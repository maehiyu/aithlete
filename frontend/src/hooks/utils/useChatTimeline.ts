
import { useMemo } from 'react';
import type { QuestionResponse, AnswerResponse } from '../../type/type';

type ChatItem = (QuestionResponse & { type: 'question' }) | (AnswerResponse & { type: 'answer' });

export function useChatTimeline(questions: QuestionResponse[] = [], answers: AnswerResponse[] = []): ChatItem[] {
	return useMemo(() => {
		const merged: ChatItem[] = [
			...questions.map(q => ({ ...q, type: 'question' as const })),
			...answers.map(a => ({ ...a, type: 'answer' as const })),
		];
		merged.sort((a, b) => new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime());
		return merged;
	}, [questions, answers]);
}
