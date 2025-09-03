import { createUser } from "./participantService";
import { createChat } from "./chatService";
import { v4 as uuidv4 } from "uuid";
import type { ParticipantResponse, ChatDetailResponse } from "../type/type";

export async function createAICoachParticipant(sport: string): Promise<ParticipantResponse> {
  return createUser({
    name: "AIコーチ",
    email: `ai_coach_${sport}@example.com`,
    role: "ai_coach",
    sports: [sport],
    iconUrl: null
  });
}

export async function createAIChatWithQuestion({
  userId,
  aiCoachSport,
  questionContent
}: {
  userId: string;
  aiCoachSport: string;
  questionContent: string;
}): Promise<{ chat: ChatDetailResponse; aiCoachParticipant: ParticipantResponse; questionId: string }> {
  const aiCoach = await createAICoachParticipant(aiCoachSport);
  const chat = await createChat({ participantIds: [userId, aiCoach.id] });
  const questionId = uuidv4();
  return { chat, aiCoachParticipant: aiCoach, questionId };
}
