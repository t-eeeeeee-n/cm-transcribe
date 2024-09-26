import React from 'react';

type TypographyProps = {
    children: React.ReactNode;
    className?: string;
};

export const Typography: React.FC<TypographyProps> = ({ children, className }) => {
    return (
        <p className={`text-base text-gray-700 ${className}`}>
            {children}
        </p>
    );
};
