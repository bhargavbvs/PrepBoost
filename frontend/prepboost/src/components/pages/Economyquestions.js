import React, { useState } from 'react';

const Economyquestions =(props) => {

	

	const questions = [
		{
			questionText: ' What do you mean by the supply of goods?',
			answerOptions: [
				{ answerText: 'Stock available for sale', isCorrect: false },
				{ answerText: ' Total stock in the warehouse', isCorrect: false },
				{ answerText: ' Quantity of the goods offered for sale at a particular price per unit of time', isCorrect: true },
				{ answerText: 'The actual production of the goods', isCorrect: false },
			],
		},
		{
			 questionText: 'What do you mean by under conditions of a perfect competition in the product market?',
			answerOptions: [
				{ answerText: 'MRP > VMP', isCorrect: false },
				{ answerText: 'MRP = VMP', isCorrect: true },
				{ answerText: 'VMP > MRP', isCorrect: false },
				{ answerText: 'None of the above', isCorrect: false },
			],
		},
		{
			questionText: ' What do you mean by a mixed economy?',
			answerOptions: [
				{ answerText: 'Public and private sectors', isCorrect: true },
				{ answerText: 'Modern and traditional industries', isCorrect: false },
				{ answerText: ' Foreign and domestic investments', isCorrect: false },
				{ answerText: 'Commercial and subsistence farming', isCorrect: false },
			],
		},
		{
			questionText: 'Which of the following is the reason for the decline in the child sex ratio in India? ',
			answerOptions: [
				{ answerText: 'None of the above', isCorrect: false },
				{ answerText: 'Incentives for a boy child from the government', isCorrect: false },
				{ answerText: 'Low fertility rate.', isCorrect: false },
				{ answerText: 'Female foeticide', isCorrect: true },
			],
		},
	];

	const data = props.data

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
							{data}
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

export default Economyquestions;