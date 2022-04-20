import React, { useEffect,useState } from 'react';

export default function TopicQuestions(props) {

	const topicid = props.topicid
	const subtopicid = props.subtopicid

	
	const handleAnswerOptionClick = (Answer,Option) => {
		if (Answer ===  Option) {
			setScore(score + 1);
		}

		const nextQuestion = currentQuestion + 1;
		if (nextQuestion < question.length) {
			setCurrentQuestion(nextQuestion);
		} else {
			setShowScore(true);
		}
	};


	const [currentQuestion, setCurrentQuestion] = useState(0);
	const [showScore, setShowScore] = useState(false);
	const [score, setScore] = useState(0);
	const [question, setQuestions] = useState([]);

	useEffect(() => {
		const loadquestions = async () => {
			let item = { topicid,subtopicid }
			console.warn("details", item)

			let result = await fetch(`http://6f49-2600-8807-c0c0-d400-cc94-a9ef-a41-1466.ngrok.io/topicwise/${topicid}/${subtopicid}`,{
				method: 'GET',
				headers: {
					"Content-Type": 'application/json',
					"Accept": 'application/json'
				}
			})
			result = await result.json()
			console.warn("result", result)
			setQuestions(result)

		}

		loadquestions();
	}, []);

	
	return (
		<div className='topicquestions'>
			{
				question[0] ? <>{showScore ? (
					<div className='score-section'>
						You scored {score} out of {question.length}
					</div>
				) : (
					<>
						
						<div className='question-section'>
							<div className='question-count'>
								<span>Question {currentQuestion + 1}</span>/{question.length}
								
							</div>
							<div className='question-text'>{question[currentQuestion].Question} </div>

							<br />
							<div className='answer-section'>
								<button onClick={() => handleAnswerOptionClick(question[currentQuestion].Answer,"A")}>{question[currentQuestion].Option1}</button>
								<button onClick={() => handleAnswerOptionClick(question[currentQuestion].Answer,"B")}>{question[currentQuestion].Option2}</button>
								<button onClick={() => handleAnswerOptionClick(question[currentQuestion].Answer,"C")}>{question[currentQuestion].Option3}</button>
								<button onClick={() => handleAnswerOptionClick(question[currentQuestion].Answer,"D")}>{question[currentQuestion].Option4}</button>
							</div>
							
							<div className='score'>
								You scored {score} out of {question.length}
							</div>
						</div>
					</>
				)}
				</> : "Loading...."
			}
		</div>
	);
}