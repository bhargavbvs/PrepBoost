import React, { useState } from 'react';

export default function Agriculturequestions() {
	const questions = [
		{
			questionText: ' Which of the following are plantation crops?',
			answerOptions: [
				{ answerText: 'tea', isCorrect: false },
				{ answerText: ' banana', isCorrect: false },
				{ answerText: ' Tea, coffee, banana and sugarcane', isCorrect: true },
				{ answerText: 'sugarcane', isCorrect: false },
			],
		},
		{
			 questionText: 'Rabi crops are:',
			answerOptions: [
				{ answerText: 'sown in winter ', isCorrect: false },
				{ answerText: 'sown in winter and harvested in summer', isCorrect: true },
				{ answerText: 'sown in summer and harvested in winter', isCorrect: false },
				{ answerText: 'None of the above', isCorrect: false },
			],
		},
		{
			questionText: 'Kharif crops are grown:',
			answerOptions: [
				{ answerText: 'with the onset of monsoon and harvested in September-October', isCorrect: true },
				{ answerText: 'with the onset of winter and harvested in summer', isCorrect: false },
				{ answerText: 'none of the above ', isCorrect: false },
				{ answerText: 'with onset of Autumn and harvested in summer', isCorrect: false },
			],
		},
		{
			questionText: 'The main food crop of Kharif season is:',
			answerOptions: [
				{ answerText: 'pulses', isCorrect: false },
				{ answerText: 'jowar', isCorrect: false },
				{ answerText: 'wheat', isCorrect: false },
				{ answerText: ' Rice', isCorrect: true },
			],
		},
	];

	const [currentQuestion, setCurrentQuestion] = useState(0);
	const [showScore, setShowScore] = useState(false);
	const [score, setScore] = useState(0);

	const handleAnswerOptionClick = (isCorrect) => {
		if (isCorrect) {
			setScore(score + 1);
		}

		const nextQuestion = currentQuestion + 1;
		if (nextQuestion < questions.length) {
			setCurrentQuestion(nextQuestion);
		} else {
			setShowScore(true);
		}
	};
	return (
    <div className='topicquestions'>
			{showScore ? (
				<div className='score-section'>
					You scored {score} out of {questions.length}
				</div>
			) : (
				<>
					<div className='question-section'>
						<div className='question-count'>
							<span>Question {currentQuestion + 1}</span>/{questions.length}
						</div>
						<div className='question-text'>{questions[currentQuestion].questionText} </div>
					
					<br/>
					<div className='answer-section'>
						{questions[currentQuestion].answerOptions.map((answerOption) => (
							<button onClick={() => handleAnswerOptionClick(answerOption.isCorrect)}>{answerOption.answerText}</button>
						))}
					</div>
					</div>
				</>
			)}
		</div>
	);
}